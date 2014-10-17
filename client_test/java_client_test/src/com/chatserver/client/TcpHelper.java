package com.chatserver.client;

import java.io.*;
import java.net.Socket;

public class TcpHelper {
	private Socket socket;
	private InputStream in;
	private OutputStream out;
	
	// ���ӷ��������ɹ�����true��ʧ�ܷ���false
	public boolean connect(String ip, int port) {
		try {
			this.socket = new Socket(ip, port);
			this.in = this.socket.getInputStream();
			this.out = this.socket.getOutputStream();
		} catch (Exception e) {
			e.printStackTrace();
			return false;
		}
		return true;
	}
	
	// �ر�����
	public void close() {
		try {
			this.socket.close();
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
	
	// ������Ϣ���ɹ�����true, ʧ�ܷ���false
	public boolean send(byte[] sendData) {
		try {
			out.write(sendData);
		} catch (Exception e) {
			e.printStackTrace();
			return false;
		}
		return true;
	}

	// ��ȡ��Ϣ���ɹ����ض�ȡ�ֽ�����ʧ�ܷ���-1
	public int read(byte[] readData) {
		int readSize = -1;
		try {
			readSize = in.read(readData);
		} catch (Exception e) {
			e.printStackTrace();
			readSize = -1;
		}
		return readSize;
	}
}
