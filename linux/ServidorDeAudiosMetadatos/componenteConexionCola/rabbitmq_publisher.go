package componenteconexioncola

import (
	"encoding/json"
	"fmt"
	"servidorquealmacenacanciones/capaFachadaServices/models"

	"github.com/streadway/amqp"
)

type RabbittPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitPublisher() (*RabbittPublisher, error) {
	conn, err := amqp.Dial("amqp://admin:1234@192.168.12.98:5672/")
	if err != nil {
		return nil, fmt.Errorf("error conectando a RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Error abriendo el canal: %v", err)
	}

	q, err := ch.QueueDeclare(
		"notificaciones_canciones",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("error declarando cola: %v", err)
	}

	return &RabbittPublisher{
		conn:    conn,
		channel: ch,
		queue:   q,
	}, nil
}

func (p *RabbittPublisher) PublicarNotificacion(msg models.NotificacionCancion) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("error convirtiendo a JSON: %v", err)
	}

	err = p.channel.Publish(
		"",
		p.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return fmt.Errorf("error publicando: %v", err)
	}

	return nil
}

func (p *RabbittPublisher) Cerrar() {
	if p.channel != nil {
		p.channel.Close()
	}
	if p.conn != nil {
		p.conn.Close()
	}
}
