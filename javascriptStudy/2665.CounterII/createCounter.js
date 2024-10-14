/**
 * @param {integer} init
 * @return { increment: Function, decrement: Function, reset: Function }
 */
var createCounter = function(init) {
  // クロージャの一部としてinitの値の現在の状態を保持しているcurrnetValueはcreatCounter外から更新はできない
  let currentValue = init;
  return {
  increment: function() {
          currentValue += 1;
          return currentValue;
      },
      decrement: function() {
          currentValue -= 1;
          return currentValue;
      },
      reset: function() {
          currentValue = init;
          return currentValue;
      }
  }
};

/**
* const counter = createCounter(5)
* counter.increment(); // 6
* counter.reset(); // 5
* counter.decrement(); // 4
*/
