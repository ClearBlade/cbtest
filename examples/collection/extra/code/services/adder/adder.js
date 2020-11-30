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

function adder(req,resp){

    ClearBlade.init({ request: req });

    var params = req.params;
    const lhs = params.lhs || 0;
    const rhs = params.rhs || 0;
    const result = lhs + rhs;

    var coll = ClearBlade.Collection({ collectionName: 'results' });

    coll.create({ lhs, rhs, result }, function(err, data) {
      if (err) {
        resp.error(data);
      } else {
        resp.success(result);
      }
    });
}
