import {Spending} from "../../../types/data-types";
import {NativeStackScreenProps} from "@react-navigation/native-stack";
import React from "react";

export type HomeStackParams = {
    Home: {}
    AddNewSpend: {setSpends:  React.Dispatch<React.SetStateAction<Spending[]>>}
    ViewSpend: {spend: Spending, setSpends: React.Dispatch<React.SetStateAction<Spending[]>>}
}

export type HomeProps = NativeStackScreenProps<HomeStackParams, 'Home'>;

