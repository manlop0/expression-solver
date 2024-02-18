"use client";
import React, { useState } from "react";

type Props = {
  submiting: boolean;
  operation: string;
  defaultValue: number;
  setFormData: React.Dispatch<React.SetStateAction<FormData[]>>;
};

type FormData = {
  name: string;
  duration: number;
};

export default function TextField({
  submiting,
  operation,
  defaultValue,
  setFormData,
}: Props) {
  const [value, setValue] = useState<string>(defaultValue.toString());

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    let result = event.target.value.replace(/\D/g, "");

    setValue(result);
    setFormData((prev: FormData[]) => {
      const newData: FormData[] = prev.map((item) => {
        if (item.name === operation) {
          return { ...item, duration: result === "" ? 0 : parseInt(result) };
        }
        return item;
      });
      return newData;
    });
  };

  return (
    <div className="relative mt-4">
      <label htmlFor={operation} className="text-black ">{`Execution time of '${operation}' operation`}</label>
      <input
        id={operation}
        disabled={submiting}
        value={value}
        onChange={handleChange}
        type="text"
        placeholder="0"
        name={operation}
        className="w-full h-12 relative text-black pl-4 border-gray-500 border-solid border rounded-md bg-transparent op-input"
      />
      <p className="absolute right-4 text-[18px] text-gray-600 top-1/2">sec</p>
    </div>
  );
}
