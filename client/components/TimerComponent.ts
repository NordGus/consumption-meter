import { humanizeTime } from "../helpers/humanizeTime.ts";

class TimerComponent extends HTMLElement {
  private timer!: HTMLElement;
  private timerId: ReturnType<typeof setInterval> | undefined;

  constructor() {
    super();
  }

  connectedCallback(): void {
    const consume = this.querySelector<HTMLButtonElement>("#consume");
    const produce = this.querySelector<HTMLButtonElement>("#produce");
    const toggle = this.querySelector<HTMLButtonElement>("#toggle");
    const tabs = this.querySelectorAll<HTMLElement>("#tab");

    this.timer = this.querySelector<HTMLElement>("#timer")!;

    if (!consume || !produce || !toggle || !tabs) return;

    this.swapOnClick(consume, "Consume", produce);
    this.swapOnClick(produce, "Produce", consume);

    toggle.addEventListener("click", () => {
      for (const tab of tabs) {
        tab.classList.toggle("flex");
        tab.classList.toggle("hidden");
      }
    });
  }

  swapOnClick(elem: HTMLButtonElement, name: string, other: HTMLButtonElement): void {
    elem.addEventListener("click", () => {
      if (elem.innerText === name) {
        this.startTimer(name);
        elem.innerText = "Stop";
        other.disabled = true;
      } else {
        this.stopTimer();
        elem.innerText = name;
        other.disabled = false;
      }
    });
  }

  startTimer(name: string): void {
    if (name === "Consume") {
      this.timer.classList.remove("text-green-600");
      this.timer.classList.add("text-red-600");
    } else {
      this.timer.classList.remove("text-red-600");
      this.timer.classList.add("text-green-600");
    }

    const start = Date.now();

    this.timerId = setInterval(() => {
      this.timer.innerText = humanizeTime(Date.now() - start);
    }, 20);
  }

  stopTimer(): void {
    clearInterval(this.timerId);
    this.timerId = undefined;
  }
}

customElements.define("timer-component", TimerComponent);
