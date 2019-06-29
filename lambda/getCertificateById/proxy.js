'use strict'
const lambda = require('./index.js')
let event = {'alias':'debug'};
let context = null
let callback = (err) => {
    console.log(err)
}
// lambdaのエントリーポイントを呼び出し
lambda.handler(event, context, callback)
