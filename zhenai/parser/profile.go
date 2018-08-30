package parser

import (
	"myProject/Single-crawler-go/engine"
	"regexp"
	"strconv"
	"myProject/Single-crawler-go/model"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var workspaceRe = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var xingzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)

func ParserProfile(contents []byte, name string) engine.ParseResult {

	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(contents, ageRe))

	if err != nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, heightRe))

	if err != nil {
		profile.Height = height
	}
	//收入
	profile.Income = extractString(contents,incomeRe)
	//婚姻状况
	profile.Marriage = extractString(contents,marriageRe)
	//教育
	profile.Education = extractString(contents,educationRe)
	//户籍
	profile.Hokou = extractString(contents,hokouRe)
	//职业
	profile.Occupation = extractString(contents,occupationRe)
	//星座
	profile.Xingzuo = extractString(contents,xingzuoRe)
	//工作地
	profile.Workplace = extractString(contents,workspaceRe)

	result := engine.ParseResult{
		Iterms:[]interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
