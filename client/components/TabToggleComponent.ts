import { TOGGLE_TAB_EVENT } from "@Helpers/constants.ts";
import { ToggleTabEvent } from "@Helpers/types.ts";

class TabToggleComponent extends HTMLButtonElement {
  private readonly target: string;

  constructor() {
    super();

    this.target = this.dataset.target as string;
  }

  connectedCallback(): void {
    this.addEventListener("click", this.onClick);
    window.addEventListener(TOGGLE_TAB_EVENT, (event) => this.onTriggerTimer(event as ToggleTabEvent), false);
  }

  disconnectedCallback(): void {
    this.removeEventListener("click", this.onClick);
    window.removeEventListener(TOGGLE_TAB_EVENT, (event) => this.onTriggerTimer(event as ToggleTabEvent), false);
  }

  onClick(): void {
    const event = new CustomEvent<{ target: string }>(TOGGLE_TAB_EVENT, {
      detail: { target: this.target },
      bubbles: true,
    });

    this.dispatchEvent(event);
  }

  onTriggerTimer(event: ToggleTabEvent): void {
    this.classList.toggle("hidden", event.detail.target === this.target);
  }
}

customElements.define("tab-toggle", TabToggleComponent, { extends: "button" });
