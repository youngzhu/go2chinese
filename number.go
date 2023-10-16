package go2chinese

const (
	ten            = 10
	shi            = ten
	hundred        = 10 * ten
	bai            = hundred
	thousand       = 10 * hundred
	qian           = thousand
	tenThousand    = 10 * thousand // 万
	wan            = tenThousand
	hundredMillion = tenThousand * tenThousand // 亿=万万
	yi             = hundredMillion

	jiao  = "角"
	fen   = "分"
	zheng = "整"
	yuan  = "元"
)

var rmbBig = []string{
	"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖",
}

var rmbUnit = map[int]string{
	ten:            "拾",
	hundred:        "佰",
	thousand:       "仟",
	tenThousand:    "万",
	hundredMillion: "亿",
}

func Number(n int64) string {
	return ""
}
