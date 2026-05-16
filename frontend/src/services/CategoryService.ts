export const GetCategory = async () => {
  const response = await fetch(`${import.meta.env.VITE_API_URL}/category/`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  const data = await response.json();
  return data;
};

export const GetCategoryById = async (id: number) => {
  const response = await fetch(
    `${import.meta.env.VITE_API_URL}/category/${id}`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    },
  );
  const data = await response.json();
  return data;
};
