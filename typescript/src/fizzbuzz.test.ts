import { beforeEach, describe, expect, test } from "@jest/globals";
import { initializeFizzBuzz } from "./fizzbuzz";

describe("FizzBuzz", (): void => {
  let container: HTMLOListElement;
  let button: HTMLButtonElement;

  beforeEach((): void => {
    container = document.createElement("ol");
    button = document.createElement("button");
    button.setAttribute("type", "button");

    initializeFizzBuzz(container, button);
  });

  test("initialized", (): void => {
    expect(container.childElementCount).toEqual(0);
  });

  test("Fizz Buzz", (): void => {
    for (let i = 1; i < 100; i++) {
      button.click();

      expect(container.childElementCount).toEqual(i);

      let expectation: string;
      if (i % 15 === 0) {
        expectation = "FizzBuzz";
      } else if (i % 5 === 0) {
        expectation = "Buzz";
      } else if (i % 3 === 0) {
        expectation = "Fizz";
      } else {
        expectation = String(i);
      }
      expect(container.lastChild?.textContent).toEqual(expectation);
    }
  });
});
