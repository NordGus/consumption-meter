import { TOGGLE_TAB_EVENT } from "@Helpers/constants.ts";
import { ToggleTabEvent } from "@Helpers/types.ts";

class TabViewComponent extends HTMLElement {
  private readonly name: string;

  constructor() {
    super();

    this.name = this.dataset.name as string;
  }

  connectedCallback(): void {
    window.addEventListener(TOGGLE_TAB_EVENT, (event) => this.onToggle(event as ToggleTabEvent));
  }

  disconnectedCallback(): void {
    window.removeEventListener(TOGGLE_TAB_EVENT, (event) => this.onToggle(event as ToggleTabEvent));
  }

  onToggle(event: ToggleTabEvent): void {
    this.classList.toggle("flex", event.detail.target === this.name);
    this.classList.toggle("hidden", event.detail.target !== this.name);
  }
}

customElements.define("tab-view", TabViewComponent);
