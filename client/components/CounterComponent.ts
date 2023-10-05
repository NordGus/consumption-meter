export default class CounterElement extends HTMLButtonElement {
  private count: number;

  constructor() {
    super();

    this.count = 0;
    this.innerHTML = `count is ${this.count}`;
  }

  connectedCallback(): void {
    this.innerHTML = `count is ${this.count}`;
    this.addEventListener("click", this.increment);
  }

  disconnectedCallback(): void {
    this.removeEventListener("click", this.increment);
  }

  increment(): void {
    this.count++;
    this.innerHTML = `count is ${this.count}`;
  }
}

customElements.define("rom-counter", CounterElement, { extends: "button" });
