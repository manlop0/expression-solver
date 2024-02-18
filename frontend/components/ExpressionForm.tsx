import React, { useState } from "react";
import axios, { AxiosError } from "axios";
import { apiUrl } from "@/utils";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faArrowAltCircleRight } from "@fortawesome/free-solid-svg-icons";

type ExpressionData = {
  value: string;
  date: Date;
};

type Props = {
  setSubmiting: React.Dispatch<React.SetStateAction<boolean>>;
  setResponseStatus: React.Dispatch<React.SetStateAction<number>>
  submiting: boolean
}

export default function ExpressionForm({submiting, setSubmiting, setResponseStatus}:Props) {
  const [value, setValue] = useState<string>("");


  const handleSubmit = async () => {
    try {
      setSubmiting(true);
      const date = new Date();
      const data: ExpressionData = {
        value: value,
        date: date,
      };
      const response = await axios.post(`${apiUrl}/api/addExpression`, data);
      setResponseStatus(response.status);
    } catch (error) {
      if (axios.isAxiosError(error)) {
        if (!error.response) {
          console.log("No Server Response");
        } else if (error.response.status === 400) {
          setResponseStatus(400)
        } else {
          setResponseStatus(500);
        }
      }
    } finally {
      setValue("");
      setSubmiting(false);
      setTimeout(() => {
        setResponseStatus(0);
      }, 3000);
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key == "Enter") {
      handleSubmit();
    }
  };

  return (
    <div className="w-full max-w-xl flex gap-4 items-center m-10 relative">
      

      <input
        className="w-full h-[48px] p-4 bg-gray-200 rounded-full text-black text-lg outline-none shadow-lg "
        placeholder="write your expression..."
        value={value}
        onChange={(e) => setValue(e.target.value)}
        onKeyDown={handleKeyDown}
        disabled={submiting}
      />
      <button
        onClick={handleSubmit}
        disabled={submiting}
        className="bg-secondary  h-[48px] w-[48px] rounded-full p-4 items-center text-lg justify-center flex text-nowrap"
      >
        <FontAwesomeIcon
          icon={faArrowAltCircleRight}
          color="#ffffff"
          fontSize={48}
        />
      </button>
    </div>
  );
}
