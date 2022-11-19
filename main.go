package main

import (
	prototype "golang-design-pattern/07_prototype"
	visitor "golang-design-pattern/23_visitor"
)

const parCount = 100


func main()  {

	//00--simplefactory mode
	//api := simplefactory.NewAPI(1)
	//fmt.Println(api.Say("Tom"))
	//api = simplefactory.NewAPI(2)
	//fmt.Println(api.Say("Tom"))
	//00--simplefactory mode


	//04--factory mode
	//var factory factorymethod.OperatorFactory
	//factory = factorymethod.PlusOperatorFactory{}
	//if factorymethod.Compute(factory, 1, 2) != 3 {
	//	fmt.Println("error with factory method pattern")
	//}
	//factory = factorymethod.MinusOperatorFactory{}
	//if factorymethod.Compute(factory, 4, 2) != 2 {
	//	fmt.Println("error with factory method pattern")
	//}
	//04--factory mode

	//05--abstractfactory
	//var factory abstractfactory.DAOFactory
	//factory = &abstractfactory.RDBDAOFactory{}
	//abstractfactory.GetMainAndDetail(factory)
	//
	//factory = &abstractfactory.XMLDAOFactory{}
	//abstractfactory.GetMainAndDetail(factory)
	//05--abstractfactory


	//03--singleton
	//inst := singleton.GetInstance()
	//inst2 := singleton.GetInstance()
	//if inst2 != inst{
	//	fmt.Println("instance not equal")
	//}else{
	//	fmt.Println("instance equal")
	//}


	//start Test  Parallel  Singleton
	//start := make(chan struct{})
	//wg := sync.WaitGroup{}
	//wg.Add(parCount)
	//instances := [parCount]singleton.Singleton{}
	//for i := 0; i < parCount; i++ {
	//	go func(index int) {
	//		// 协程阻塞，等待channel被关闭才能继续运行
	//		<-start
	//		instances[index] = singleton.GetInstance()
	//		fmt.Println("i=",i)
	//		wg.Done()
	//	}(i)
	//}
	//// 关闭channel，所有协程同时开始运行，实现并行(parallel)
	//close(start)
	//wg.Wait()
	//for i := 1; i < parCount; i++ {
	//	if instances[i] != instances[i-1] {
	//		fmt.Println("instance is not equal")
	//	}else{
	//		fmt.Println("instance equal")
	//	}
	//}
	//03--singleton


	//06 --builder
	//b := &builder.Builder1{}
	//director := builder.NewDirector(b)
	//director.Construct()
	//res := b.GetResult()
	//if res != "123"{
	//	fmt.Printf("Builder1 fail expect 123 acture %s", res)
	//}
	//fmt.Println("res=",res)
	//
	//
	//b2 := &builder.Builder2{}
	//director2 := builder.NewDirector(b2)
	//director2.Construct()
	//res2 := b2.GetResult()
	//if res2 != 6 {
	//	fmt.Println("Builder2 fail expect 6 acture %d", res)
	//}
	//fmt.Println("res2=",res2)
	//06-- builder

	//07原型模式
	//manager := prototype.NewPrototypeManager()
	//t1 := &Type1{
	//	name: "type1",
	//}
	//manager.Set("t1", t1)
	//tget := manager.Get("t1")
	//t2 := tget.Clone()
	//if tget == t2 {
	//	fmt.Println("error! get clone not working")
	//}
	//fmt.Println("tget=", t2.Clone())
	//
	////Test Clone From Manager
	//c := manager.Get("t1").Clone()
	//t_clone := c.(*Type1)
	//if t_clone.name != "type1" {
	//	fmt.Println("error:",t_clone.name)
	//}else{
	//	fmt.Println("t_clone.name=",t_clone.name)
	//}
	//07原型模式

	//02 adapter
	//adap := adapter.NewAdaptee()
	//target := adapter.NewAdapter(adap)
	//res := target.Request()
	//if res != "adaptee method"{
	//	fmt.Println("res=",res)
	//}
	//fmt.Println("111,res=",res)
	//02 adapter

	//20decorator
	//var c decorator.Component
	//c  = &decorator.ConcreteComponent{}
	//c = decorator.WrapAddDecorator(c, 10)
	//c = decorator.WrapMulDecorator(c, 8)
	//res := c.Calc()
	//
	//fmt.Printf("res %d\n", res)
	//20decorator


	//10--observer
	//suj := observer.NewSubject()
	//red1 := observer.NewReader("reader1")
	//red2 := observer.NewReader("reader2")
	//red3 := observer.NewReader("reader3")
	//suj.Attach(red1)
	//suj.Attach(red2)
	//suj.Attach(red3)
	//suj.UpdateContext("observer mode")
	//10--observer

	//23--访问者模式
	c := &visitor.CustomerCol{}
	c.Add(visitor.NewEnterpriseCustomer("A company"))
	c.Add(visitor.NewEnterpriseCustomer("B company"))
	c.Add(visitor.NewIndividualCustomer("bob"))
	c.Accept(&visitor.ServiceRequestVisitor{})

	c1 := &visitor.CustomerCol{}
	c1.Add(visitor.NewEnterpriseCustomer("A company"))
	c1.Add(visitor.NewIndividualCustomer("bob"))
	c1.Add(visitor.NewEnterpriseCustomer("B company"))
	c1.Accept(&visitor.AnalysisVisitor{})
	//23--访问者模式



}


type Type1 struct {
	name string
}

func (t *Type1) Clone() prototype.Cloneable {
	tc := *t
	return &tc
}

