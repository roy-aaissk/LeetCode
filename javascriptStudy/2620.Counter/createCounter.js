/**
 * @param {number} n
 * @return {Function} counter
 */
// クロージャと呼ばれるもので8行目の無名関数のnの変数の参照先は6行目のnの値になるそのため内包している関数から外側の関数に結果を返すことができる
var createCounter = function(n) {
    
  return function() {
      return n++
  };
};

/** 
* const counter = createCounter(10)
* counter() // 10
* counter() // 11
* counter() // 12
*/
