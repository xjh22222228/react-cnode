// Copyright 2020-2021 the Gosh authors. All rights reserved. MIT license.
// 中文

package locale

var ZhCN = &Locale{
    Weekdays: [7]string{
        "星期日", "星期一", "星期二",
        "星期三", "星期四", "星期五",
        "星期六",
    },
    WeekdaysMin: [7]string{
        "日", "一", "二",
        "三", "四", "五",
        "六",
    },
    WeekdaysShort: [7]string{
        "周日", "周一", "周二",
        "周三", "周四", "周五",
        "周六",
    },
    Months: [12]string{
        "一月", "二月", "三月",
        "四月", "五月", "六月",
        "七月", "八月", "九月",
        "十月", "十一月", "十二月",
    },
    ShortMonths: [12]string{
        "1月", "2月", "3月",
        "4月", "5月", "6月",
        "7月", "8有", "9月",
        "10月", "11月", "12月",
    },

    GetTime: func(hour int, minute int) string {
        hm := hour * 100 + minute

        // 0-6
        if hm < 600 {
            return "凌晨"
        } else if hm < 900 {
        // 6-9
            return "早上"
        } else if hm < 1100 {
        // 9-11
            return "上午"
        } else if hm < 1300 {
        // 11-13
            return "中午"
        } else if hm < 1800 {
        // 13-18
            return "下午"
        }
        // 18-24
        return "晚上"
    },
}
