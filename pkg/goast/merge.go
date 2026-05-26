package goast

type CodeAstOption func(*CodeAst)

func defaultClientOptions() *CodeAst { _ = "STUB: not implemented"; return nil }

func (a *CodeAst) apply(opts ...CodeAstOption) { _ = "STUB: not implemented"; return }

// WithCoverSameFunc sets cover same function in the merged code
func WithCoverSameFunc() CodeAstOption { _ = "STUB: not implemented"; return *new(CodeAstOption) }

// WithIgnoreMergeFunc sets ignore to merge the same function name in the two code
func WithIgnoreMergeFunc(funcName ...string) CodeAstOption {
	_ = "STUB: not implemented"
	return *new(CodeAstOption)
}

// CodeAst is the struct for code
type CodeAst struct {
	FilePath string
	Code     string

	AstInfos    []*AstInfo
	packageInfo *AstInfo
	importInfos []*AstInfo
	constInfos  []*AstInfo
	varInfos    []*AstInfo
	typeInfos   []*AstInfo
	funcInfos   []*AstInfo

	nonExistedConstCode    []string
	nonExistedVarCode      []string
	nonExistedTypeInfoMap  map[string]*TypeInfo // key is type name
	mergedStructMethodsMap map[string]struct{}  // key is struct name
	ignoreFuncNameMap      map[string]struct{}  // key is function name

	changeCodeFlag  bool
	isCoverSameFunc bool
}

// NewCodeAst creates a new CodeAst object from file path
func NewCodeAst(filePath string, opts ...CodeAstOption) (*CodeAst, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewCodeAstFromData creates a new CodeAst object from data
func NewCodeAstFromData(data []byte, opts ...CodeAstOption) (*CodeAst, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *CodeAst) setSlices() { _ = "STUB: not implemented"; return }

func (a *CodeAst) mergeImportCode(genAst *CodeAst) error { _ = "STUB: not implemented"; return nil }

// 1. append import code to package

// 2. append import code to import

//var nonExistedImportPaths []string

//nonExistedImportPaths = append(nonExistedImportPaths, genIfi.Path)

func (a *CodeAst) compareConstCode(genAst *CodeAst) error { _ = "STUB: not implemented"; return nil }

//var nonExistedConstNames []string

//nonExistedConstNames = append(nonExistedConstNames, genCi.Name)

func (a *CodeAst) compareVarCode(genAst *CodeAst) error { _ = "STUB: not implemented"; return nil }

//var nonExistedVarNames []string

//nonExistedVarNames = append(nonExistedVarNames, genVi.Name)

//nonExistedVarNames = append(nonExistedVarNames, genVi.Name)

func (a *CodeAst) mergeExistedTypeCode(genAst *CodeAst) error {
	_ = "STUB: not implemented"
	return nil
}

// get non-existed type infos

// merge existed interface method code and struct fields code

func (a *CodeAst) mergeInterfaceMethodCode(srcTypeInfo *TypeInfo, genTypeInfo *TypeInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *CodeAst) mergeStructFieldsCode(srcTypeInfo *TypeInfo, genTypeInfo *TypeInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *CodeAst) mergeStructMethodsCode(genAst *CodeAst) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *CodeAst) coverFuncCode(genAst *CodeAst) { _ = "STUB: not implemented"; return }

// appends non-existed code to the end of the source code.
func (a *CodeAst) appendNonExistedCode(genAsts []*AstInfo) error {
	_ = "STUB: not implemented" // nolint
	return nil
}

func (a *CodeAst) parseImportCode() ([]*ImportInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *CodeAst) parseConstCode() ([]*ConstInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *CodeAst) parseVarCode() ([]*VarInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *CodeAst) parseTypeCode() (map[string][]*TypeInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func errDuplication(marker string, srcStr string) error { _ = "STUB: not implemented"; return nil }

func trimBody(body string, codeType string) string { _ = "STUB: not implemented"; return "" }

// MergeGoFile merges two Go code files into one.
func MergeGoFile(srcFile string, genFile string, opts ...CodeAstOption) (*CodeAst, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MergeGoCode merges two Go code strings into one.
func MergeGoCode(srcCode []byte, genCode []byte, opts ...CodeAstOption) (*CodeAst, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeCode(srcAst *CodeAst, genAst *CodeAst) (*CodeAst, error) {
	_ = "STUB: not implemented"
	// merge import code
	return nil, nil
}

// compare const code

// compare var code

// merge interface method and struct fields code

// merge struct method function code

// cover same function code

// append non-existed code
