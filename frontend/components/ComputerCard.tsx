import React from "react";
import Loader from "./Loader";
import { faPause } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

type WorkersData = {
  id: number;
  working: boolean;
  workingon: string;
};

export default function ComputerCard({ data }: { data: WorkersData }) {
  return (
    <div className="flex min-h-[100px] h-fit w-full items-center py-2 px-12 bg-secondary rounded-xl gap-4 shadow-lg">
      {data.working ? <Loader /> : <FontAwesomeIcon icon={faPause} fontSize={60} />}
      <div className="text-center ml-8">
        <h1 className="text-xl ">{`Computer server №${data.id}`}</h1>
        <h2 className="text-base text-gray-200 ">
          {data.working ? `Working on "${data.workingon}"` : "Waiting for work"}
        </h2>
      </div>
    </div>
  );
}
