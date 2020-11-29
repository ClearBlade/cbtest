/**
 * Type: Micro Service
 * Description: A short-lived service which is expected to complete within a fixed period of time.
 * @param {CbServer.BasicReq} req
 * @param {string} req.systemKey
 * @param {string} req.systemSecret
 * @param {string} req.userEmail
 * @param {string} req.userid
 * @param {string} req.userToken
 * @param {boolean} req.isLogging
 * @param {[id: string]} req.params
 * @param {CbServer.Resp} resp
 */

function faker(req,resp){
    // These are parameters passed into the code service
    var params = req.params;

    resp.success({
        foo: "foo",
        bar: "bar",
        baz: "baz",
        one: 1,
        two: 2,
        three: 3,
        random: Math.random() * 100,
        message: "The quick brown fox jumps over the lazy dog",
    })
}
