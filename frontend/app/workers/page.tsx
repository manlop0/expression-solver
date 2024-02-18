"use client";

import ComputerCard from "@/components/ComputerCard";
import Spinner from "@/components/Spinner";
import { apiUrl } from "@/utils";
import axios from "axios";
import React, { useEffect, useState } from "react";

type WorkersData = {
  id: number;
  working: boolean;
  workingon: string;
};

export default function WorkersPage() {
  const [data, setData] = useState<WorkersData[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setIsLoading(true);
        const response = await axios.get(`${apiUrl}/api/getWorkers`);
        const sortedData: WorkersData[] = response.data.sort(
          (a: WorkersData, b: WorkersData) => a.id - b.id
        );
        setData(sortedData);
      } catch (error) {
        console.log("Error while getting workers data: ", error);
      } finally {
        setIsLoading(false);
      }
    };
    fetchData();
  }, []);

  return (
    <div>
      {isLoading ? (
        <div className="flex mt-4 w-full h-full justify-center items-center">
          <Spinner />
        </div>
      ) : (
        <div className="justify-center flex mt-8 w-full">
          <div className="items-center flex flex-col gap-4 w-full max-w-[1200px] mx-4">
            {data.map((el) => (
              <ComputerCard key={`Worker${el.id}`} data={el} />
            ))}
          </div>
        </div>
      )}
    </div>
  );
}
