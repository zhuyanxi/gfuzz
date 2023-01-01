package oracle

import oraclert "github.com/zhuyanxi/gfuzz/pkg/oraclert"

func TestHello() {
	oracleEntry := oraclert.BeforeRun()
	defer oraclert.AfterRun(oracleEntry)
	println("hello")
}
