// TimerTrigger Types
export type TimerTriggerAction = "consume" | "produce";
export type TimerTriggerType = "start" | "stop";
export type TriggerTimerEvent = CustomEvent<{ action: "consume" | "produce"; trigger: "start" | "stop" }>;

// TabToggle Types
export type ToggleTabEvent = CustomEvent<{ target: string }>;
