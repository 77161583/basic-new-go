function IDValidator() {
    var param = {
        error: {
            longNumber: '长数字存在精度问题，请使用字符串传值！ Long number is not allowed, because the precision of the Number In JavaScript.',
        },
    }
    var util = {
        checkArg: function (id) {
            var argType = typeof id

            switch (argType) {
                case 'number':
                    // long number not allowed
                    id = id.toString()
                    if (id.length > 15) {
                        this.error(param.error.longNumber)
                        return false
                    }
                    break
                case 'string':
                    break
                default:
                    return false
            }
            id = id.toUpperCase()
            var code = null
            if (id.length === 18) {
                // 18位
                code = {
                    body: id.slice(0, 17),
                    checkBit: id.slice(-1),
                    type: 18,
                }
            } else if (id.length === 15) {
                // 15位
                code = {
                    body: id,
                    type: 15,
                }
            } else {
                return false
            }
            return code
        },
        // 地址码检查
        checkAddr: function (addr) {
            const arr = [
                11, 12, 13, 14, 15,
                21, 22, 23,
                31, 32, 33, 34, 35, 36, 37,
                41, 42, 43, 44, 45, 46,
                50, 51, 52, 53, 54,
                61, 62, 63, 64, 65,
                71, 81, 82,
            ]
            return arr.includes(addr.slice(0, 2) * 1)
        },
        // 生日码检查
        checkBirth: function (birth) {
            var year, month, day
            if (birth.length == 8) {
                year = parseInt(birth.slice(0, 4), 10)
                month = parseInt(birth.slice(4, 6), 10)
                day = parseInt(birth.slice(-2), 10)
            } else if (birth.length == 6) {
                year = parseInt('19' + birth.slice(0, 2), 10)
                month = parseInt(birth.slice(2, 4), 10)
                day = parseInt(birth.slice(-2), 10)
            } else {
                return false
            }
            // TODO 是否需要判断年份
            /*
             * if( year<1800 ){ return false; }
             */
            // TODO 按月份检测
            if (month > 12 || month === 0 || day > 31 || day === 0) {
                return false
            }

            return true
        },
        // 顺序码检查
        checkOrder: function (order) {
            // 暂无需检测

            return true
        },
        // 加权
        weight: function (t) {
            return Math.pow(2, t - 1) % 11
        },
        // 随机整数
        rand: function (max, min) {
            min = min || 1
            return Math.round(Math.random() * (max - min)) + min
        },
        // 数字补位
        str_pad: function (str, len, chr, right) {
            str = str.toString()
            len = len || 2
            chr = chr || '0'
            right = right || false
            if (str.length >= len) {
                return str
            } else {
                for (var i = 0, j = len - str.length; i < j; i++) {
                    if (right) {
                        str = str + chr
                    } else {
                        str = chr + str
                    }
                }
                return str
            }
        },
        // 抛错
        error: function (msg) {
            var e = new Error()
            e.message = 'IDValidator: ' + msg
            throw e
        },
    }
    var _IDValidator = function () {
        // 建立cache
        this.cache = {}
    }
    _IDValidator.prototype = {
        isValid: function (id) {
            var code = util.checkArg(id)
            if (code === false) {
                return false
            }
            // 查询cache
            if (this.cache.hasOwnProperty(id) && typeof this.cache[id].valid !== 'undefined') {
                return this.cache[id].valid
            } else {
                if (!this.cache.hasOwnProperty(id)) {
                    this.cache[id] = {}
                }
            }

            var addr = code.body.slice(0, 6)
            var birth = code.type === 18 ? code.body.slice(6, 14) : code.body.slice(6, 12)
            var order = code.body.slice(-3)

            if (!(util.checkAddr(addr) && util.checkBirth(birth) && util.checkOrder(order))) {
                this.cache[id].valid = false
                return false
            }

            // 15位不含校验码，到此已结束
            if (code.type === 15) {
                this.cache[id].valid = true
                return true
            }

            /* 校验位部分 */

            // 位置加权
            var posWeight = []
            for (var i = 18; i > 1; i--) {
                var wei = util.weight(i)
                posWeight[i] = wei
            }

            // 累加body部分与位置加权的积
            var bodySum = 0
            var bodyArr = code.body.split('')
            for (var j = 0; j < bodyArr.length; j++) {
                bodySum += parseInt(bodyArr[j], 10) * posWeight[18 - j]
            }

            // 得出校验码
            var checkBit = 12 - (bodySum % 11)
            if (checkBit == 10) {
                checkBit = 'X'
            } else if (checkBit > 10) {
                checkBit = checkBit % 11
            }
            checkBit = typeof checkBit === 'number' ? checkBit.toString() : checkBit

            // 检查校验码
            if (checkBit !== code.checkBit) {
                this.cache[id].valid = false
                return false
            } else {
                this.cache[id].valid = true
                return true
            }
        },
    } // _IDValidator
    return new _IDValidator()
}

export default IDValidator
