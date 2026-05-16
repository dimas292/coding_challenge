import { DATE_FORMAT_OPTIONS } from "./constants";

export const formatDate = (date?: string | null): string => {
  if (!date) return "-";
  const d = new Date(date);
  return isNaN(d.getTime()) ? "-" : d.toLocaleDateString("en-US", DATE_FORMAT_OPTIONS);
};

export const compareDates = (date1?: string | null, date2?: string | null): number => {
  const d1 = date1 ? new Date(date1).getTime() : 0;
  const d2 = date2 ? new Date(date2).getTime() : 0;
  return d1 - d2;
};
