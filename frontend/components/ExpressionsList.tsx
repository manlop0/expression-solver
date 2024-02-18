
import React, { useEffect, useState } from "react";
import ExpressionCard from "./ExpressionCard";
import axios from "axios";
import Spinner from "./Spinner";
import { apiUrl } from "@/utils";

type ExpressionData = {
  id: number;
  value: string;
  date: Date;
  status: number;
  result: string;
};

export default function ExpressionsList({submiting}:{submiting:boolean}) {
  const [data, setData] = useState<ExpressionData[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setIsLoading(true);
        const response = await axios.get(`${apiUrl}/api/getExpressions`);
        setData(response?.data.reverse());
      } catch (error) {
        console.log(error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchData();
  }, [submiting]);

  return (
    <div>
      {isLoading ? (
        <div className="flex justify-center items-center w-full h-full">
          <Spinner />
        </div>
      ) : (
        <div className="flex flex-col gap-8">
          {data.map((el) => (
            <ExpressionCard key={`exp${el.id}`} data={el}/>
          ))}
        </div>
      )}
    </div>
  );
}
