package com.chatserver.client;

//---------------------------���ݰ�����------------------------------------

//��1���ֽ��򣺴��ģʽ

//��2�����ݰ���ɣ����� + ���� + ����

//������4�ֽڣ�int���������ݰ��ĳ���

//���ͣ�4�ֽڣ�int

//���壺�ֽ����飬byte[]

//���������������Ĵ��䣬�����ɽṹ��protobuf���л����ٽ���AES���ܵõ���

//����ֻ������

//--------------------------------------------------------------------------

public class PacketHelper {
	
	// ���ݰ������ͣ� ���pb.proto
	public static int PK_ClientLogin = 1;
	public static int PK_ServerAcceptLogin = 2;
	public static int PK_ClientLogout = 3;
	public static int PK_ClientPing = 4;
	public static int PK_C2CTextChat = 5;
	
	// ���ݰ����
	private int pacLen;
	private int pacType;
	private byte[] pacData;
	
	// typeΪ�������ͣ� dataΪprotobuf���л����Ѿ������˵�����
	public PacketHelper(int type, byte[] data) {
		this.pacData = data;
		this.pacLen = 8 + this.pacData.length;
		this.pacType = type;
	}
	
	public byte[] getBytes() {
		byte[] bytes = new byte[this.pacLen];
		System.arraycopy(ConvHelper.intToBytes(this.pacLen), 0, bytes, 0, 4);
		System.arraycopy(ConvHelper.intToBytes(this.pacType), 0, bytes, 4, 4);
		System.arraycopy(this.pacData, 0, bytes, 8, this.pacData.length);
		return bytes;
	}
	
	public static void main(String[] args) {
		// ����
		PacketHelper ph = new PacketHelper(PacketHelper.PK_ClientLogin, "hello world, ����дjava".getBytes());
		byte[] b = ph.getBytes();
		
		byte[] len = new byte[4];
		byte[] type = new byte[4];
		byte[] data = new byte[b.length-8];
		System.arraycopy(b, 0, len, 0, 4);
		System.arraycopy(b, 4, type, 0, 4);
		System.arraycopy(b, 8, data, 0, b.length-8);
		
		System.out.println(ConvHelper.bytesToInt(len));
		System.out.println(ConvHelper.bytesToInt(type));
		System.out.println(new String(data));
	}
}
