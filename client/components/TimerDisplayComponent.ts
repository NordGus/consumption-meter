import { TRIGGER_TIMER_EVENT } from "@Helpers/constants.ts";
import { humanizeTime } from "@Helpers/humanizeTime.ts";
import { TriggerTimerEvent } from "@Helpers/types.ts";

class TimerDisplayComponent extends HTMLElement {
  private timerId: ReturnType<typeof setInterval> | undefined;

  constructor() {
    super();
  }

  connectedCallback(): void {
    window.addEventListener(TRIGGER_TIMER_EVENT, (event) => this.trigger(event as TriggerTimerEvent), false);
  }

  disconnectedCallback(): void {
    window.removeEventListener(TRIGGER_TIMER_EVENT, (event) => this.trigger(event as TriggerTimerEvent), false);
  }

  private trigger(event: TriggerTimerEvent): void {
    if (event.detail.trigger === "stop") return this.stopTimer();
    return this.startTimer(event.detail.action);
  }

  startTimer(name: string): void {
    const start = Date.now();

    this.classList.toggle("text-green-600", name === "produce");
    this.classList.toggle("text-red-600", name === "consume");

    this.timerId = setInterval(() => {
      this.innerText = humanizeTime(Date.now() - start);
    }, 20);
  }

  stopTimer(): void {
    clearInterval(this.timerId);
    this.timerId = undefined;
  }
}

customElements.define("timer-display", TimerDisplayComponent);
