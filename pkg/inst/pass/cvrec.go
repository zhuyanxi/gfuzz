package pass

import (
	"gfuzz/pkg/inst"
	"go/ast"
	"go/token"
	"strconv"

	"golang.org/x/tools/go/ast/astutil"
)

// CvResPass, Conditional Variable Pass.
type CvRecPass struct{}

func (p *CvRecPass) Name() string {
	return "cvrec"
}

func (p *CvRecPass) Run(iCtx *inst.InstContext) error {
	inst.AddImport(iCtx.FS, iCtx.AstFile, "gooracle", "gooracle")
	iCtx.AstFile = astutil.Apply(iCtx.AstFile, instChOps, nil).(*ast.File)
	return nil
}

func (p *CvRecPass) Deps() []string {
	return nil
}

func instCvOps(c *astutil.Cursor) bool {
	switch concrete := c.Node().(type) {
	case *ast.ExprStmt:
		if callExpr, ok := concrete.X.(*ast.CallExpr); ok {
			if selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok { // like `mu.Lock()`
				var matched bool = true
				var op string
				switch selectorExpr.Sel.Name {
				case "Broadcast":
					op = "Broadcast"
				case "Signal":
					op = "Signal"
				case "Wait":
					op = "Wait"
				default:
					matched = false
				}

				if matched {
					intID := int(Uint16OpID)
					newCall := NewArgCallExpr("gooracle", "StoreOpInfo", []ast.Expr{&ast.BasicLit{
						ValuePos: 0,
						Kind:     token.STRING,
						Value:    "\"" + op + "\"",
					}, &ast.BasicLit{
						ValuePos: 0,
						Kind:     token.INT,
						Value:    strconv.Itoa(intID),
					}})
					c.InsertBefore(newCall)
					boolNeedInstrument = true // We need to import gooracle
					records = append(records, strconv.Itoa(intID)+":trad"+op)
					Uint16OpID++
				}
			}

		}
	}
	return true
}
