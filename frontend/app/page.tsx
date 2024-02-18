'use client'
import ExpressionForm from "@/components/ExpressionForm";
import ExpressionsList from "@/components/ExpressionsList";
import { useState } from "react";

export default function Home() {

  const [submiting, setSubmiting] = useState<boolean>(false);
  const [responseStatus, setResponseStatus] = useState<number>(0);

  return (
    <div className="items-center flex m-8 flex-col relative">
      {responseStatus !== 0 && (
        <div
          className={`absolute z-10 w-fit top-2 left-0 right-0 ml-auto mr-auto flex items-center justify-center rounded-md p-4 shadow-lg ${
            responseStatus === 200
              ? "bg-green-600"
              : responseStatus === 400
              ? "bg-red-500"
              : "bg-black border-2 border-white"
          }`}
        >
          <p className="text-xl text-center">
            {responseStatus === 200
              ? "Expression successfully accepted"
              : responseStatus === 400
              ? "Invalid expression"
              : "Some problems on Back-end"}
          </p>
        </div>
      )}

      <ExpressionForm setResponseStatus={setResponseStatus} setSubmiting={setSubmiting} submiting={submiting}/>
      <ExpressionsList submiting={submiting}/>
    </div>
  );
}
