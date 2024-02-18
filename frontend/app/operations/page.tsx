"use client";
import Spinner from "@/components/Spinner";
import TextField from "@/components/TextField";
import { apiUrl } from "@/utils";
import axios from "axios";
import React, { useEffect, useState } from "react";

type OperationData = {
  name: string;
  duration: number;
};


export default function OperationsPage() {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [data, setData] = useState<OperationData[]>([]);
  const [formData, setFormData] = useState<OperationData[]>(data)
  const [submiting, setSubmiting] = useState<boolean>(false)

  useEffect(() => {
    const fetchData = async () => {
      try {
        setIsLoading(true);
        const response = await axios.get(`${apiUrl}/api/getOperations`);
        setData(response.data);
        setFormData(response.data)
      } catch (error) {
        console.log("Error while getting operations data: ", error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchData();
  }, []);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    try {
      setSubmiting(true)
      const response = await axios.put(`${apiUrl}/api/changeOperations`, formData)
      setData(response.data)
    } catch (error) {
      console.log("Error while patching the Data: ", error)
    } finally {
      setTimeout(() => {
        setSubmiting(false)
      }, 1000);
    }
  };

  return (
    <div className="mt-16 flex justify-center">
      {isLoading ? (
        <Spinner />
      ) : (
        <form
          onSubmit={handleSubmit}
          className=" bg-white border-blue-500 border-2 items-center shadow-xl bg-opacity-80 flex py-8 px-24 max-w-[600px] w-full flex-col"
        >
          {data.map((el) => (
            <TextField submiting={submiting} key={el.name} defaultValue={el.duration} operation={el.name} setFormData={setFormData}/>
          ))}
          <button disabled={submiting}  type="submit" className={`w-24 mt-8 p-2 rounded-md shadow-xl ${submiting ? "bg-primary" : "bg-secondary"}`}>
            Apply
          </button>
        </form>
      )}
    </div>
  );
}
