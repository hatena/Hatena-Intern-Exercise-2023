import { initializeFizzBuzz } from "./fizzbuzz";

const container = document.getElementById("fizzbuzz") as HTMLOListElement;
const button = document.getElementById("do") as HTMLButtonElement;
initializeFizzBuzz(container, button);
