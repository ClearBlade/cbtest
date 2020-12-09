/**
 * Type: Stream Service
 * Description: A service that does not have an execution timeout which allows for infinite execution of logic.
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

function streamLogger(req, resp) {
  ClearBlade.init({ request: req });

  const TOPIC = "$share/group/messages";
  const COLLECTION = "messages";
  var messaging = ClearBlade.Messaging();
  messaging.subscribe(TOPIC, WaitLoop);

  function debug(topic, payload) {
    messaging.publish(topic, payload);
  }

  function WaitLoop(err, data) {

    if (err) {
      debug('error', "Subscribe failed: " + data);
      resp.error(data);
    }

    debug('success', "Subscribed to Shared Topic. Starting Loop.");

    while (true) {
      messaging.waitForMessage([TOPIC], function(err, msg, topic) {
        if (err) {
          const msg = "Failed to wait for message: " + err + " " + msg + "  " + topic;
          debug('error', msg);
          resp.error(msg);
        }
        processMessage(msg, topic);
      });
    }
  }

  function processMessage(msg, topic) {

    const data = JSON.parse(msg);
    const coll = ClearBlade.Collection({ collectionName: COLLECTION });
    coll.create(data, function(err, data) {

      if (err) {
        const msg = "Failed to put message in collection: " + COLLECTION + " " + data;
        debug('error', msg);
        resp.error(msg);
      }

      debug('success', "Message processed " + msg + " " + topic);
    });
  }
}
