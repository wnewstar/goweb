/**
 * 业务无关函数
 */

package tool

import (
    "sort"
    "encoding/json"
)

/**
 * @desc    获取签名数据信息
 * @name    GetSignStr
 * @date    2020-03-07
 * @author  wnewstar
 * @param   p               interface{}                     签名数据
 * @return  back            string                          签名数据
 */
func GetSignStr(p interface{}) (back string) {
    ks := []string{}
    data := p.(map[string]string)
    for k := range data {
        ks = append(ks, k)
    }
    sort.Strings(ks)
    for _, k := range ks { back = back + k + "=" + data[k] }

    return
}

/**
 * @desc    获取签名数据映射
 * @name    GetSignMap
 * @date    2020-03-07
 * @author  wnewstar
 * @param   p               interface{}                     原始参数
 * @param   f               string                          父级键名
 * @return  back            map[string]string{}             签名映射
 */
func GetSignMap(p interface{}, f string) (back map[string]string) {
    back = make(map[string]string)
    switch p.(type) {
        case []interface{}:
            t := p.([]interface{})
            s := t[len(t) - 1]
            for k, v := range GetSignMap(s, f) { back[k] = v }
        case map[string]interface{}:
            for k, v := range p.(map[string]interface{}) {
                if v == nil {
                    continue
                }
                switch v.(type) {
                    case string:
                        s := v.(string)
                        if len(s) > 0 { back[f + k] = s + "&" }
                    case bool, float64:
                        s, err := json.Marshal(v)
                        if err == nil { back[f + k] = string(s) + "&" }
                    case []interface{}:
                        t := v.([]interface{})
                        s := t[len(t) - 1]
                        for sk, sv := range GetSignMap(s, f + k) { back[sk] = sv }
                    case map[string]interface{}:
                        for sk, sv := range GetSignMap(v, f + k) { back[sk] = sv }
                }
            }
    }

    return
}
