package genappevent

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func FmtImports(pkgs []string) string {
	if len(pkgs) == 0 {
		return ""
	}

	groups := make([][]string, 2)

	for _, pkg := range pkgs {
		if len(strings.Split(pkg, "/")) < 3 && !strings.Contains(pkg, ".") {
			groups[0] = append(groups[0], pkg)
			continue
		}
		groups[1] = append(groups[1], pkg)
	}

	b := new(bytes.Buffer)
	for _, group := range groups {
		group := group
		sort.Slice(group, func(i, j int) bool {
			return group[i] < group[j]
		})
		for _, pkg := range group {
			_, err := b.WriteString(strconv.Quote(pkg))
			if err != nil {
				panic(err)
			}
			_, err = b.WriteRune('\n')
			if err != nil {
				panic(err)
			}
		}
		_, err := b.WriteRune('\n')
		if err != nil {
			panic(err)
		}
	}

	return fmt.Sprintf(`import (
%s
		)`,
		b.String(),
	)
}

func ToUpperCamel(s string) string {
	if s == "" {
		return s
	}
	firstNotLowerIndex := strings.IndexFunc(s, func(c rune) bool {
		return !unicode.IsLower(c)
	})
	if firstNotLowerIndex == -1 {
		firstNotLowerIndex = len(s)
	}
	if commonInitialisms[s[:firstNotLowerIndex]] {
		return strings.ToUpper(s[:firstNotLowerIndex]) + s[firstNotLowerIndex:]
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func ToLowerCamel(s string) string {
	if s == "" {
		return s
	}
	firstNotUpperIndex := strings.IndexFunc(s, func(c rune) bool {
		return !unicode.IsUpper(c)
	})
	if firstNotUpperIndex == -1 {
		firstNotUpperIndex = len(s)
	}
	if commonInitialisms[s[:firstNotUpperIndex]] {
		return strings.ToLower(s[:firstNotUpperIndex]) + s[firstNotUpperIndex:]
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// from https://github.com/golang/lint
var commonInitialisms = map[string]bool{
	"acl":   true,
	"api":   true,
	"ascii": true,
	"cpu":   true,
	"css":   true,
	"dns":   true,
	"eof":   true,
	"guid":  true,
	"html":  true,
	"http":  true,
	"https": true,
	"id":    true,
	"ip":    true,
	"json":  true,
	"lhs":   true,
	"qps":   true,
	"ram":   true,
	"rhs":   true,
	"rpc":   true,
	"sla":   true,
	"smtp":  true,
	"sql":   true,
	"ssh":   true,
	"tcp":   true,
	"tls":   true,
	"ttl":   true,
	"udp":   true,
	"ui":    true,
	"uid":   true,
	"uuid":  true,
	"uri":   true,
	"url":   true,
	"utf8":  true,
	"vm":    true,
	"xml":   true,
	"xmpp":  true,
	"xsrf":  true,
	"xss":   true,
}
