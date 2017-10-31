# go-db-helper
an lightweight orm framework


## install
> go get -u github.com/aidonggua/go-db-helper

## Getting Started

**first**
connect database

    err := dh.Connect("127.0.0.1", 3306, "root", "root", "db_helper")

**then**
register your struct to map the specified table

	err := dh.Register("tb_test", &TestRegistry{})

**finally**
you can add database records by calling the **Insert** method or by calling the **InsertSelective** method

	id, err := dh.Insert(&TestRegistry{Name: "melon", Age: 18, CreateTime: time.Now()})

	id, err := dh.InsertSelective(&TestRegistry{Name: "aa", Age: 18})

> note: the different bettwen Insert and InsertSelective is that the InsertSelective will ignore the zero field


 you can update the specified database record by calling the **UpdateById** method or by calling the **UpdateByIdSelective** method

## author
**melon**

- e.yehaoo@gmail.com
