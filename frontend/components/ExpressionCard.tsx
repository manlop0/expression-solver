import React from "react";
import Queued from "./Queued";
import Spinner from "./Spinner";
import { faCheckCircle } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

type Data = {
  id: number;
  value: string;
  date: Date;
  status: number;
  result: string;
};

export default function ExpressionCard({ data }: { data: Data }) {
  return (
    <div
      className="flex max-w-xl w-fit items-center min-h-[100px] min-w-[320px] h-fit py-2 px-6 bg-secondary rounded-xl gap-4 shadow-lg"
    >
      {data.status === 0 ? (
        <Queued />
      ) : data.status === 1 ? (
        <Spinner />
      ) : (
        <FontAwesomeIcon icon={faCheckCircle} fontSize={48} color="#ffffff"/>
      )}
      <div className="flex flex-col w-full items-center">
        <div className=" text-xl">{`${data.value}=${data.result ? data.result : "?"}`}</div>
        <div className="text-gray-200 text-base">
          {data.status === 0 ? (
            <p>in Queue...</p>
          ) : data.status === 1 ? (
            <p>is Solving...</p>
          ) : (
            <p>Solved</p>
          )}
        </div>
        <p className="text-gray-300 text-sm">{`CreatedAt: ${data.date.toLocaleDateString()} ${data.date.toLocaleTimeString()}`}</p>
      </div>
    </div>
  );
}
