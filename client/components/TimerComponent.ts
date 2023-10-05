function humanTime(ms: number): string {
  const mils = ms % 1000;
  const out = [];

  ms = Math.floor(ms / 1000);

  if (ms > 0 && ms % 60 < 10) {
    out.push(`0${ms % 60}.${mils}`);
  } else if (ms > 0) {
    out.push(`${ms % 60}.${mils}`);
  } else {
    out.push(`00.${mils}`);
  }

  ms = Math.floor(ms / 60);

  if (ms > 0 && ms % 60 < 10) {
    out.push(`0${ms % 60}`);
  } else if (ms > 0) {
    out.push(`${ms % 60}`);
  }

  ms = Math.floor(ms / 60);

  if (ms > 0 && ms < 10) {
    out.push(`0${ms}`);
  } else if (ms > 0) {
    out.push(`${ms}`);
  }

  return out.reverse().join(":");
}

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

    const view = this.querySelector<HTMLButtonElement>("#view");
    const timings = this.querySelector<HTMLButtonElement>("#timings");
    this.timer = this.querySelector<HTMLElement>("#timer")!;

    if (!consume || !produce || !toggle || !timings || !view) return;

    this.swapOnClick(consume, "Consume", produce);
    this.swapOnClick(produce, "Produce", consume);

    toggle.addEventListener("click", () => {
      this.timer.classList.toggle("hidden");
      timings.classList.toggle("hidden");
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
      this.timer.innerText = humanTime(Date.now() - start);
    }, 20);
  }

  stopTimer(): void {
    clearInterval(this.timerId);
    this.timerId = undefined;
  }
}

customElements.define("timer-component", TimerComponent);
