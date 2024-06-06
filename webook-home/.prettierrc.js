module.exports = {
    // 不使用缩进符，使用空格缩进
    useTabs: false,
    // 缩进4个空格
    tabWidth: 4,
    // 每行最多160个字符
    printWidth: 160,
    // 行尾不必有分号
    semi: false,
    // 使用单引号
    singleQuote: true,
    // 箭头函数，只有一个参数的时候，也需要括号
    arrowParens: 'always',
    // 大括号内的首尾需要空格
    bracketSpacing: true,
    // 使用拖尾逗号（主要缓解增加一行对象属性，导致git变更记录是两行的情况）
    trailingComma: 'es5'
}