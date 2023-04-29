import {Spending} from "../../../types/data-types";
import {NativeStackScreenProps} from "@react-navigation/native-stack";

export type HomeStackParams = {
    Home: {}
    AddNewSpend: {}
    ViewSpend: {spend: Spending}
}

export type HomeProps = NativeStackScreenProps<HomeStackParams, 'Home'>;

