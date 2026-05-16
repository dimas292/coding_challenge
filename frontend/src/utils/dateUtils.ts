import { DATE_FORMAT_OPTIONS } from "./constants";

export const formatDate = (date: string): string => {
  return new Date(date).toLocaleDateString("en-US", DATE_FORMAT_OPTIONS);
};

export const compareDates = (date1: string, date2: string): number => {
  return new Date(date1).getTime() - new Date(date2).getTime();
};
