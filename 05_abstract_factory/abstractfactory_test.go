package abstractfactory

//func getMainAndDetail(factory DAOFactory) {
//	factory.CreateOrderMainDAO().SaveOrderMain()
//	factory.CreateOrderDetailDAO().SaveOrderDetail()
//}

func ExampleRdbFactory() {
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	GetMainAndDetail(factory)
	// Output:
	// rdb main save
	// rdb detail save
}

func ExampleXmlFactory() {
	var factory DAOFactory
	factory = &XMLDAOFactory{}
	GetMainAndDetail(factory)
	// Output:
	// xml main save
	// xml detail save
}
