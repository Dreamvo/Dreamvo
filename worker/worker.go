package worker

import (
	"fmt"
	"github.com/dreamvo/gilfoyle/config"
	"github.com/dreamvo/gilfoyle/ent"
	"github.com/dreamvo/gilfoyle/logging"
	"github.com/dreamvo/gilfoyle/storage"
	"github.com/dreamvo/gilfoyle/transcoding"
	"github.com/streadway/amqp"
)

const (
	HlsVideoEncodingQueue   string = "HlsVideoEncoding"
	EncodingEntrypointQueue string = "EncodingEntrypoint"
	EncodingFinalizerQueue  string = "EncodingFinalizer"
)

type AMQPClient interface {
	Channel() (*amqp.Channel, error)
	Close() error
	IsClosed() bool
}

type Channel interface {
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error)
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
}

type Queue struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
	Handler    func(*Worker, amqp.Delivery)
}

var queues = []Queue{
	{
		Name:       HlsVideoEncodingQueue,
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
		Handler:    hlsVideoEncodingConsumer,
	},
	{
		Name:       EncodingFinalizerQueue,
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
		Handler:    encodingFinalizerConsumer,
	},
	{
		Name:       EncodingEntrypointQueue,
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
		Handler:    encodingEntrypointConsumer,
	},
}

type Options struct {
	Host        string
	Port        int
	Username    string
	Password    string
	Concurrency uint
	Logger      logging.ILogger
	Storage     storage.Storage
	Database    *ent.Client
	Transcoder  transcoding.ITranscoder
	Config      config.Config
}

type Worker struct {
	Queues      map[string]amqp.Queue
	Client      AMQPClient
	logger      logging.ILogger
	concurrency uint
	storage     storage.Storage
	dbClient    *ent.Client
	transcoder  transcoding.ITranscoder
	config      config.Config
}

func New(opts Options) (*Worker, error) {
	conn, err := amqp.Dial(fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		opts.Username,
		opts.Password,
		opts.Host,
		opts.Port,
	))
	if err != nil {
		return nil, err
	}

	return &Worker{
		Queues:      map[string]amqp.Queue{},
		Client:      conn,
		logger:      opts.Logger,
		concurrency: opts.Concurrency,
		storage:     opts.Storage,
		dbClient:    opts.Database,
		transcoder:  opts.Transcoder,
		config:      opts.Config,
	}, nil
}

func (w *Worker) Init() error {
	ch, err := w.Client.Channel()
	if err != nil {
		return err
	}

	for _, q := range queues {
		queue, err := ch.QueueDeclare(
			q.Name,       // name
			q.Durable,    // durable
			q.AutoDelete, // delete when unused
			q.Exclusive,  // exclusive
			q.NoWait,     // no-wait
			q.Args,       // arguments
		)
		if err != nil {
			return err
		}

		w.Queues[q.Name] = queue
	}

	return nil
}

func (w *Worker) Consume() error {
	ch, err := w.Client.Channel()
	if err != nil {
		return fmt.Errorf("error creating channel: %e", err)
	}

	err = ch.Qos(
		int(w.concurrency),
		0,
		false,
	)
	if err != nil {
		return fmt.Errorf("error setting prefetch option (concurrency): %e", err)
	}

	for _, q := range queues {
		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			false,  // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			map[string]interface{}{},
		)
		if err != nil {
			return fmt.Errorf("error consuming %s queue: %e", q.Name, err)
		}

		// Start goroutines to consume messages
		go func(q Queue) {
			for d := range msgs {
				go q.Handler(w, d)
			}
		}(q)
	}

	return nil
}

func (w *Worker) Close() error {
	err := w.Client.Close()
	if err != nil {
		return err
	}

	return nil
}
