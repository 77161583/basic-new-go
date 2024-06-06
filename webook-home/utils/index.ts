export function formatDate(value: any, fmt = 'yyyy-MM-dd hh:mm:ss') {
  if (!value || value === '0000-00-00 00:00:00') return ''
  let val = value.toString()
  val = val.length === 10 ? val * 1000 : val.length === 13 ? val * 1 : val
  const dateObj = new Date(val)
  let result = fmt
  const obj = {
    'M+': dateObj.getMonth() + 1,
    'd+': dateObj.getDate(),
    'h+': dateObj.getHours(),
    'm+': dateObj.getMinutes(),
    's+': dateObj.getSeconds(),
    'q+': Math.floor((dateObj.getMonth() + 3) / 3),
    S: dateObj.getMilliseconds(),
  }
  if (/(y+)/.test(fmt)) {
    const matchRes = fmt.match(/(y+)/) || []
    const fmt_year = matchRes[0]
    result = result.replace(fmt_year, (dateObj.getFullYear() + '').substring(4 - fmt_year.length))
  }
  for (const key in obj) {
    const rule = new RegExp('(' + key + ')')
    if (rule.test(result)) {
      const matchRes = fmt.match(rule) || []
      const fmt_matched = matchRes[0]
      result = result.replace(fmt_matched, fmt_matched.length === 1 ? obj[key] : ('00' + obj[key]).substring(('' + obj[key]).length))
    }
  }
  return result
}

export function listToTree(list = [], id = 'id', pid = 'pid', children = 'children') {
  const data = list.map((item: any) => ({ ...item }))
  const hash = {}
  data.forEach((item) => {
    hash[item[id]] = item
  })
  const result: any = []
  data.forEach((item) => {
    const parent = hash[item[pid]]
    if (parent) {
      ;(parent[children] || (parent[children] = [])).push(item)
    } else {
      result.push(item)
    }
  })
  return result
}

export function delaySomeTime(ms: number) {
  return new Promise(resolve => setTimeout(resolve, ms))
}

/**
 * 将str中的转义字符还原成html字符
 * @see UE.utils.unhtml(String);
 * @method html
 * @param { String } str 需要逆转义的字符串
 * @return { String } 逆转义后的字符串
 * @example
 * ```javascript
 *
 * var str = '&lt;body&gt;&amp;&lt;/body&gt;';
 *
 * //output: <body>&</body>
 * console.log( UE.utils.html( str ) );
 *
 * ```
 */
export function str2html(str: string) {
  str = str
    ? str.replace(/&amp;/g, function(m: any) {
      return {
        '&amp;': '&',
      }[m]
    })
    : ''
  return str
    ? str.replace(/&(gt|lt|quot|amp|#39|nbsp);/g, function(m: any) {
      return {
        '&lt;': '<',
        '&amp;': '&',
        '&quot;': '"',
        '&gt;': '>',
        '&#39;': "'",
        '&nbsp;': ' '
      }[m]
    })
    : ''
}

/**
 * 拼接课程详情的url
 * @param courseId 课程或课程包id
 * @param planId 计划id
 * @param ifLive 是否是直播
 * @param ifPackage 是否是课程包
 * @returns
 */
export function formatCourseDetailUrl(courseId: string, planId: string, ifLive: boolean, ifPackage = false) {
  const paramsStr = ifPackage === true ? '&isPackage=true' : ''
  if (ifLive === true) {
    return `/live/${courseId}?planId=${planId}${paramsStr}`
  } else {
    return `/course/${courseId}?planId=${planId}${paramsStr}`
  }
}

/**
 * 格式化视频时长
 * @param sec 时长秒数
 * @returns
 */
export function formatVideoTime(sec: number) {
  sec = sec || 0
  const second = Math.floor(sec % 60)
  const minute = Math.floor((sec / 60) % 60)
  const hour = Math.floor((sec / 3600) % 24)
  const res = (minute < 10 ? '0' + minute : minute) + ':' + (second < 10 ? '0' + second : second)
  return hour > 0 ? (hour < 10 ? '0' + hour : hour) + ':' + res : res
}

/**
 * 生成随机的字符串
 * @param length
 * @returns
 */
export function generateRandomString(length:any) {
  const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
  let result = ''
  const charactersLength = characters.length
  for (let i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength))
  }
  return result
}

/**
 *
 * @param text 截取字符串
 * @param maxLength
 * @returns
 */
export function truncateText(text, maxLength) {
  if (text.length > maxLength) {
    return text.slice(0, maxLength) + '...'
  }
  return text
}

