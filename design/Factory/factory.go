package Factory

import "fmt"

type Restaurant interface {
	GetFood()
}

func NewRestaurant(name string) Restaurant {
	switch name {
	case "d":
		return &Donglaishun{}
	case "q":
		return &Qingfeng{}
	}
	return nil
}


type Donglaishun struct {
}
func (d Donglaishun) GetFood() {
	fmt.Println("东来顺的饭菜已经生成完毕就绪....")
}


type Qingfeng struct {
}
func (q Qingfeng) GetFood() {
	fmt.Println("庆丰的饭菜已经生成完毕就绪....")
}
