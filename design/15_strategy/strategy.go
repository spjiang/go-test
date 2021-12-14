package strategy

import "fmt"

// 策略模式（Strategy）

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

type PaymentContext struct {
	Name, CardID string
	Money        int
}

func NewPayment(name, cardid string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardid,
			Money:  money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

type PaymentStrategy interface {
	Pay(*PaymentContext)
}

type Cash struct{}

func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}

type Bank struct{}

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s", ctx.Money, ctx.Name, ctx.CardID)

}
