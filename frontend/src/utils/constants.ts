// Priority filter options
export const PRIORITY_OPTIONS = [
  { text: "Low", value: "low" },
  { text: "Medium", value: "medium" },
  { text: "High", value: "high" },
] as const;

// Completed/Status filter options
export const COMPLETED_OPTIONS = [
  { text: "Done", value: true },
  { text: "Pending", value: false },
] as const;

// Priority to color mapping
export const PRIORITY_COLOR_MAP: Record<string, string> = {
  high: "red",
  medium: "orange",
  low: "green",
};

// Completed/Status to display mapping
export const COMPLETED_DISPLAY_MAP: Record<string, string> = {
  true: "Done",
  false: "Pending",
};

// Completed to color mapping
export const COMPLETED_COLOR_MAP: Record<string, string> = {
  true: "green",
  false: "default",
};

// Date format options
export const DATE_FORMAT_OPTIONS: Intl.DateTimeFormatOptions = {
  year: "numeric",
  month: "short",
  day: "numeric",
};

// Form mode constants
export const FORM_MODE = {
  CREATE: "create",
  EDIT: "edit",
} as const;

export type FormMode = (typeof FORM_MODE)[keyof typeof FORM_MODE];
