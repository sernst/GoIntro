var Worker = require('webworker-threads').Worker;

function nowMilliseconds() {
  return (new Date()).getTime();
}

function compute(lifetime) {
  var startTime = nowMilliseconds(),
      elapsedTime = 0;

  while (elapsedTime < lifetime) {
    elapsedTime = nowMilliseconds() - startTime;
  }
}

function log(index, lifetime, referenceTime) {
  var now = nowMilliseconds(),
      startTime = now - referenceTime - lifetime,
      endTime = now - referenceTime;

  console.log([
    '[' + index + ']: (' +
    (0.001*startTime).toFixed(2) + 's, ' +
    (0.001*endTime).toFixed(2) + 's)'
  ].join('\n'));
}

function createWorkers(count) {
  var workers = [];

  while (workers.length < count) {
    workers.push(new Worker(function () {

      function nowMilliseconds() {
        return (new Date()).getTime();
      }

      function compute(lifetime) {
        var startTime = nowMilliseconds(),
            elapsedTime = 0;

        while (elapsedTime < lifetime) {
          elapsedTime = nowMilliseconds() - startTime;
        }
      }

      this.onmessage = function (event) {
        compute(event.data.lifetime);
        postMessage(event.data);
        self.close();
      };
    }));
  }

  return workers;
}

function main2() {
  var referenceTime = nowMilliseconds(),
      lifetimes = [5000, 4000, 3000, 2000, 1000],
      finished = 0;

  return new Promise(function (resolve) {
    createWorkers(lifetimes.length).forEach(function (worker, index) {
      worker.onmessage = function (event) {
        log(event.data.index, event.data.lifetime, referenceTime);

        finished += 1;
        if (finished === lifetimes.length) {
          resolve();
        }
      };

      worker.postMessage({index:index, lifetime:lifetimes[index]});
    });
  })
  .then(function () {
    console.log('Done');
  });
}

function main() {
  var referenceTime = nowMilliseconds(),
      lifetimes = [5000, 4000, 3000, 2000, 1000];

  return Promise.all(
    lifetimes.map(function (lifetime, index) {
      return new Promise(function (resolve) {
        compute(lifetime);
        log(index, lifetime, referenceTime);
        resolve();
      });
  }))
  .then(function () {
    console.log('Done');
  });
}
main().then(function () {
  return main2();
});

/*
 [0]: (0.00s, 5.00s)
 [1]: (5.02s, 9.02s)
 [2]: (9.02s, 12.02s)
 [3]: (12.02s, 14.02s)
 [4]: (14.02s, 15.02s)
 Done
 [4]: (0.11s, 1.11s)
 [3]: (0.09s, 2.09s)
 [2]: (0.12s, 3.12s)
 [1]: (0.11s, 4.11s)
 [0]: (0.10s, 5.10s)
 Done
 */