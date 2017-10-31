package dh

import (
	"testing"
	"time"
	"fmt"
)

func TestConnect(t *testing.T) {
	DebugOn()
	err := Connect("127.0.0.1", 3306, "root", "root", "db_helper")
	if err != nil {
		t.Fatal(err)
	}
}

type TestRegistry struct {
	Id         int
	Name       string
	Age        int
	Sex        int
	CreateTime time.Time
}

func TestRegister(t *testing.T) {
	DebugOn()
	err := Register("tb_test", &TestRegistry{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateSql(t *testing.T) {
	TestRegister(t)
	_, _, err := generateInsertSql(&TestRegistry{})
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = generateUpdateByIdSql(&TestRegistry{Id: 1})
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = generateInsertSelectiveSql(&TestRegistry{Name: "aa", Age: 18})
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = generateUpdateByIdSelectiveSql(&TestRegistry{Id: 2, Name: "aa", Age: 18})
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = generateDeleteByIdSql(&TestRegistry{Id: 2, Name: "aa", Age: 18})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCRUD(t *testing.T) {
	TestConnect(t)
	TestRegister(t)

	id, err := Insert(&TestRegistry{Name: "aa", Age: 18, CreateTime: time.Now()})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Sprintf("id = %d", id))

	id, err = InsertSelective(&TestRegistry{Name: "aa", Age: 18})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Sprintf("id = %d", id))

	r, err := DeleteById(&TestRegistry{Id: int(id)})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Sprintf("number of effective rows = %d", r))

	r, err = UpdateById(&TestRegistry{Id: 1, Name: "aa", Age: 18, CreateTime: time.Now()})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Sprintf("number of effective rows = %d", r))

	r, err = UpdateByIdSelective(&TestRegistry{Id: 1, Name: "rrr"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Sprintf("number of effective rows = %d", r))

	ts := []TestRegistry{}
	err = WhereEqual("name", "aa").AndEqual("age", 18).AndIsNotNull("sex").Select(&ts)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ts)
}

func TestTx(t *testing.T) {
	TestConnect(t)
	TestRegister(t)
	err := Tx(func() {
		_, err := UpdateById(&TestRegistry{Id: 1, Name: "dads", Age: 232323, CreateTime: time.Now()})
		if err != nil {
			t.Fatal(err)
		}
		panic(fmt.Errorf("manual"))
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSelectStrings(t *testing.T) {
	TestConnect(t)
	TestRegister(t)

	names := []string{}
	err := SelectStrings("select name from tb_test", &names)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(names)

	name := ""
	err = SelectString("select create_time from tb_test where id = 1", &name)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(name)
}

func TestSelectStruct(t *testing.T) {
	TestConnect(t)
	TestRegister(t)

	ts := []TestRegistry{}
	err := SelectStructs("select * from tb_test where id < 3", &ts)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ts)

}
