const html = (env, headers) => {
  return `
    <!DOCTYPE html>
    <html>
      <head>
      </head>
      <body>
        <div>
          <h3>${env}</h3>
        </div>
      </body>
    </html>
  `;
};

exports.handler = async function(event, context) {
  try {
    // const body = await html(
    //   process.env.GOOGLE_APPLICATION_CREDENTIALS
    // );
    return { 
      headers: {
        /* Required for CORS support to work */
        "Access-Control-Allow-Origin" : "*",
        /* Required for cookies, authorization headers with HTTPS */
        "Access-Control-Allow-Credentials" : true
      },
      statusCode: 200, 
      body: JSON.parse(process.env.GOOGLE_APPLICATION_CREDENTIALS).private_key.replace(/\\n/g, '\n')
      };
  } catch (err) {
    return { statusCode: 500, body: err.toString() };
  }
};
