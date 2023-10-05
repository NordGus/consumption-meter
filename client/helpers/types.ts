// TimerTrigger Types
export type TimerTriggerAction = "consume" | "create";
export type TimerTriggerType = "start" | "stop";
export type TriggerTimerEvent = CustomEvent<{ action: "consume" | "create"; trigger: "start" | "stop" }>;

// TabToggle Types
export type ToggleTabEvent = CustomEvent<{ target: string }>;
