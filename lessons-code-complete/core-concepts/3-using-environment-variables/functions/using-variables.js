exports.handler = async (event, context) => {
  console.log('process.env.GOOGLE_APPLICATION_CREDENTIALS', process.env.GOOGLE_APPLICATION_CREDENTIALS)
  const response = {
    statusCode: 200,
    body: JSON.stringify({
      headers: {
        /* Required for CORS support to work */
        "Access-Control-Allow-Origin" : "*",
        /* Required for cookies, authorization headers with HTTPS */
        "Access-Control-Allow-Credentials" : true
      },
      message: `hello ${process.env.GOOGLE_APPLICATION_CREDENTIALS}`
    }),
  }
  return callback(null, response)
}
