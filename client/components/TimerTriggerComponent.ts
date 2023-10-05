import { TRIGGER_TIMER_EVENT } from "@Helpers/constants.ts";
import { TimerTriggerAction, TimerTriggerType, TriggerTimerEvent } from "@Helpers/types.ts";

class TimerTriggerComponent extends HTMLButtonElement {
  private readonly triggerType: TimerTriggerAction;
  private active: boolean;

  constructor() {
    super();

    this.triggerType = this.dataset.action as TimerTriggerAction;
    this.active = false;
  }

  connectedCallback(): void {
    this.addEventListener("click", this.onClick);
    window.addEventListener(TRIGGER_TIMER_EVENT, (event) => this.onTriggerTimer(event as TriggerTimerEvent), false);
  }

  disconnectedCallback(): void {
    this.removeEventListener("click", this.onClick);
    window.removeEventListener(TRIGGER_TIMER_EVENT, (event) => this.onTriggerTimer(event as TriggerTimerEvent), false);
  }

  onClick(): void {
    this.active = !this.active;

    const event = new CustomEvent<{ action: TimerTriggerAction; trigger: TimerTriggerType }>(TRIGGER_TIMER_EVENT, {
      detail: { action: this.triggerType, trigger: this.active ? "start" : "stop" },
      bubbles: true,
    });

    this.dispatchEvent(event);
  }

  onTriggerTimer(event: TriggerTimerEvent): void {
    this.disabled = event.detail.action !== this.triggerType;
  }
}

customElements.define("timer-trigger", TimerTriggerComponent, { extends: "button" });
