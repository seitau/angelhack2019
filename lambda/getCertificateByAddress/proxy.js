'use strict'
const lambda = require('./index.js')
let event = { address: '0xa43e00a4d376d14117e7c97bfb57b54409f9b2b4' };
let context = null
let callback = (err) => {
    console.log(err)
}
lambda.handler(event, context, callback)
