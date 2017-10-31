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

	err := dh.Register("tb_test", &TestDbHelper{})

**finally**

you can add database records by calling the **Insert** function or by calling the **InsertSelective** function

	id, err := dh.Insert(&TestDbHelper{Name: "melon", Age: 18, CreateTime: time.Now()})

	id, err := dh.InsertSelective(&TestDbHelper{Name: "aa", Age: 18})

 you can update a specified record by calling the **UpdateById** function or by calling the **UpdateByIdSelective** function

 	n, err = dh.UpdateById(&TestDbHelper{Id: 1, Name: "aa", Age: 18, CreateTime: time.Now()})

 	n, err = dh.UpdateByIdSelective(&TestDbHelper{Id: 1, Name: "rrr"})


> note: the different bettwen Insert/UpdateById function and InsertSelective/UpdateByIdSelective function is that the 'Selective' function will ignore the zero field

you can delete a specified record by calling the **DeleteById** function

	n, err := dh.DeleteById(&TestDbHelper{Id: 1})

you can search for datas by add conditions and calling the **Select** function

	err = dh.WhereEqual("name", "melon").AndEqual("age", 18).AndIsNotNull("sex").Select(&ts)

you can pass a bussiness function into Tx function for starting a transaction

	err := df.Tx(func() error{
			_, err := UpdateById(&TestRegistry{Id: 1, Name: "melon", Age: 1, CreateTime: time.Now()})
			panic(fmt.Errorf("test rollback"))
		})

Tx function recieve a bussiness function which return an error,if the bussiness function's value is not nil, transaction will be rolled back

when you want to see the sql statement,you can call the **DebugOn** function to show the logs,if not,call the **DebugOff** function to hide the logs

	2017/10/31 15:12:53 SQL: SELECT * FROM tb_test WHERE name = ? AND age = ? AND sex IS NOT NULL
	2017/10/31 15:12:53 args: [aa 18]

## custom settings

if you want to change the timeout,you can call the SetTimeout function

**SetTimeout(timeout time.Duration)**

if you implemented the Connetion interface by yourself,you can call SetConnection function to change the framework internal Connection

**func SetConnection(c Connection)**


## author
**melon**

- e.yehaoo@gmail.com
