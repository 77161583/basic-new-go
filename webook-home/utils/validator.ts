export const isCellphone = (val: string) => {
    return /^1[3-9]\d{9}$/.test(val)
}

export const isLandline = (val: string) => {
    return /^((0\d{2,3}-\d{7,8})|(1[3456789]\d{9}))$/.test(val)
}

export const isSocialCode = (val: string) => {
    return /^[0-9A-HJ-NPQRTUWXY]{2}\d{6}[0-9A-HJ-NPQRTUWXY]{10}$/.test(val)
}

export const cellphoneValidator = (rule: any, value: string, callback: Function) => {
    value = (value || '').trim()
    if (!value) {
        callback(new Error('请填写手机号码'))
    } else if (!isCellphone(value)) {
        callback(new Error('请填写正确的手机号码'))
    } else {
        callback()
    }
}

export const landlineValidator = (rule: any, value: string, callback: Function) => {
    value = (value || '').trim()
    if (!value) {
        callback(new Error('请填写座机号码'))
    } else if (!isLandline(value)) {
        callback(new Error('请填写正确的座机号码'))
    } else {
        callback()
    }
}
