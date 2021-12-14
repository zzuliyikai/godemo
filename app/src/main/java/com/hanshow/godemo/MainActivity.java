package com.hanshow.godemo;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.widget.Toast;

import callGo.CallGo;
import go.Seq;

public class MainActivity extends AppCompatActivity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        //这里调用通过jni调用go，然后go调用MainActivity 的showToast，展示出Toast的message
        findViewById(R.id.button).setOnClickListener(view -> {
            Seq.setContext(MainActivity.this);
            CallGo.showToast();
        });
    }

    public void showToast(String message) {
        runOnUiThread(() -> Toast.makeText(MainActivity.this, message, Toast.LENGTH_SHORT).show());
    }


}