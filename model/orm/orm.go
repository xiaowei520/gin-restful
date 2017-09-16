//package orm
//
//import (
//	_ "github.com/go-sql-driver/mysql"
//	"database/sql"
//	"fmt"
//	"log"
//)
//
//func Open() *sql.DB {
//
//	db, err := sql.Open("mysql", "root:@tcp(10.211.55.9:3306)/dqcenter?charset=utf8")
//	defer db.Close() //遇到错误-连接数据库关闭
//	if err != nil {
//		log.Fatalln(err)
//		fmt.Println(err.Error())
//	}
//	fmt.Println(db)
//	log.Println(db)
//	return db
//}ß

package orm

import (
	"github.com/opentracing/opentracing-go"
	"log"
)

/**
简单的interface 继承
 */
type context interface {
	OpenTracingSpan() opentracing.Span
}

type Model struct {
	Context context
}

/**
传递参数Model 类型
 */
func (m Model) Trace() opentracing.Span {
	if m.Context != nil {
		if span := m.Context.OpenTracingSpan(); span != nil {
			comp := "orm"
			s := opentracing.StartSpan(comp+":GetUserById", opentracing.ChildOf(span.Context()))
			s.SetTag("component", comp)
			s.SetTag("span.kind", "server")

			return s
		} else {
			log.Println("trace faile context span nil")
		}
	} else {
		log.Println("trace faile orm context nil")
	}
	return nil
}
